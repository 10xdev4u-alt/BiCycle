# 🚴‍♂️ **SVCE HOSTEL CYCLE ACCESS SYSTEM - GUARD APP INTEGRATED MODEL**

---

## 📋 **EXECUTIVE SUMMARY**

A **production-ready, guard-verified bicycle booking system** featuring **dynamic QR tickets**, **mandatory photo documentation**, and **dual-app architecture** (Student + Guard). Eliminates all spoofing vectors through **encrypted static QRs**, **time-bound dynamic tickets**, and **human-in-the-loop verification** - all at **zero hardware cost**.

---

## 🎯 **SYSTEM ARCHITECTURE**

```
┌─────────────────────────────────────────────────────────────┐
│              STUDENT MOBILE APP (React Native)              │
│  • Encrypted QR Scanner                                     │
│  • Photo Capture (Camera-Only)                             │
│  • Dynamic Ticket Generation (Pickup/Return)               │
└──────────────────────┬──────────────────────────────────────┘
                       │ HTTPS + WebSocket
                       ▼
┌─────────────────────────────────────────────────────────────┐
│              GUARD APP (React Native / PWA)                 │
│  • Dynamic QR Ticket Scanner                               │
│  • Photo Comparison (Before/After)                         │
│  • Unlock Approval Workflow                                │
└──────────────────────┬──────────────────────────────────────┘
                       │ HTTPS + WebSocket
                       ▼
┌─────────────────────────────────────────────────────────────┐
│              API GATEWAY (Node.js/Fastify)                  │
│  ┌──────────────┐  ┌──────────────┐  ┌─────────────────┐ │
│  │ Auth Service│  │ Booking Svc  │  │ Photo Svc       │ │
│  │ (JWT+OAuth) │  │ (Crypto)     │  │ (Local/S3)      │ │
│  └──────────────┘  └──────────────┘  └─────────────────┘ │
│  ┌──────────────┐  ┌──────────────┐  ┌─────────────────┐ │
│  │ Guard Svc   │  │ Ticket Svc   │  │ Admin Svc       │ │
│  │ (Approval)  │  │ (Dynamic QR) │  │ (Dashboard)     │ │
│  └──────────────┘  └──────────────┘  └─────────────────┘ │
└──────────────────────┬──────────────────────────────────────┘
                       │ 
        ┌──────────────┼──────────────┐
        ▼              ▼              ▼
┌─────────────┐  ┌──────────┐  ┌──────────────┐
│ PostgreSQL  │  │ Redis    │  │ File System  │
│ (Bookings)  │  (Real-time│  │ (Photos)     │
│ & Crypto)   │  State)    │  └──────────────┘
└─────────────┘  └──────────┘
```

**Physical Infrastructure:**
- **20 Bike QRs**: Static, encrypted, laminated
- **3 Stand QRs**: Static, encrypted (BH, GH, MG)
- **Security Guards**: Android/iOS phone with Guard App
- **Physical Locks**: Existing key locks (unchanged)

---

## 🔐 **SECURITY MODEL - 8 LAYERS**

### **Layer 1: AES-256 Encrypted Static QRs**
```javascript
// Bike QR Payload (encrypted)
{
  type: "bike",
  id: "bike_svce_07",
  initial_stand: "BH01",
  timestamp: 1699123456000
}
// Stand QR Payload (encrypted)
{
  type: "stand",
  id: "stand_bh_01",
  location: "boys_hostel",
  timestamp: 1699123456000
}
```
- **Key**: 32-byte server-side secret (stored in `.env`)
- **Advantage**: Cannot be forged; photocopy useless without decryption

### **Layer 2: College Google OAuth**
```javascript
// Login Flow
User → Google OAuth → Server verifies hd: "svce.ac.in" → Extract reg_no
→ Generate device_fingerprint → Issue JWT (24h expiry)
```
- **Device Fingerprint**: `SHA256(Phone_UUID + App_Install_ID)`
- **Binding**: JWT valid only on registered device

### **Layer 3: Redis Distributed Locks**
```javascript
// Atomic Booking Pattern
WATCH bike:status:bike_svce_07
GET bike:status:bike_svce_07  // Must be 'available'
MULTI
  SET bike:status:bike_svce_07 "booked"
  SET bike:booked_by:bike_svce_07 "2023CS0113"
  SETEX bike:expiry:bike_svce_07 300 "expired"
  SET user:active:2023CS0113 "bike_svce_07"
EXEC
```
- **Race Condition**: Only 1st request succeeds; 2nd gets error
- **Expiry**: Auto-release after 5 min (no-show handling)

### **Layer 4: Mandatory Photo Proof**
- **Pickup**: Camera-only capture (blocked gallery access)
- **Return**: Camera-only at stand location
- **EXIF Validation**: Timestamp must be within 30s of upload
- **Storage**: `pickups/{booking_id}.jpg`, `returns/{booking_id}.jpg`

### **Layer 5: Dynamic QR Tickets**
```javascript
// Pickup Ticket (JWT, expires 5 min)
{
  booking_id: "bk_7f3a9c2e",
  type: "pickup",
  bike_uuid: "bike_svce_07",
  reg_no: "2023CS0113",
  guard_code: "3847",
  exp: 1699123756
}

// Return Ticket (JWT, expires 10 min)
{
  booking_id: "bk_7f3a9c2e",
  type: "return",
  bike_uuid: "bike_svce_07",
  reg_no: "2023CS0113",
  duration_min: 45,
  exp: 1699124356
}
```
- **Guard Scans**: Verifies signature, shows student details + photos

### **Layer 6: Guard App Verification**
- **Separate Auth**: Guard JWT (issued by admin)
- **Scan Ticket**: Decodes JWT → Shows:
  - Student Name, Reg No, Photo
  - Bike ID, Pickup/Return Photos
  - Guard Code (for manual verification)
- **Approval**: Tap "✅ Approve" → Unlocks bike in system

### **Layer 7: Violation Automation**
```javascript
// No-Show Detection (Cron every 30s)
if (booking.status === 'booked' && now > booked_until) {
  db.users.no_show_count += 1
  if (no_show_count >= 3) user.is_banned = true
  releaseBike()
}

// Fast Return Detection
if (ride_duration < 2 minutes) {
  addViolation('fast_return')
}

// Late Return Detection
if (ride_duration > 120 minutes) {
  addViolation('late_return')
}
```

### **Layer 8: Immutable Audit Trail**
- **PostgreSQL**: All bookings, violations, approvals logged
- **Photos**: Never deleted, stored by booking ID
- **Redis Logs**: Temp state + guard actions backed to DB

---

## 🎯 **DETAILED USER FLOW**

### **PHASE 1: BOOKING (Student App)**

**Step 1: Scan Stand QR**
```
Student: Opens App → [SCAN STAND QR]
App: Scans → Decrypts → Sends to Server

POST /api/v1/stand/scan
{
  encrypted_stand_qr: "U2FsdGVkX1..."
}

Server Response:
{
  stand_id: "stand_bh_01",
  stand_name: "Boys Hostel - Zone A",
  message: "Stand verified. Now scan a bike."
}

App State: current_stand = "stand_bh_01"
```

**Step 2: Scan Bike QR**
```
Student: Points camera at bike QR
App: Scans → Decrypts → Sends booking request

POST /api/v1/cycles/scan-book
{
  encrypted_bike_qr: "U2FsdGVkX1...",
  current_stand: "stand_bh_01"
}

Server Checks:
├─ Decrypt bike QR → bike_uuid = "bike_svce_07"
├─ Verify bike.current_stand === "stand_bh_01"
├─ Check bike.status === 'available' (Redis)
├─ Verify user has no active booking
└─ ATOMIC: Book bike in Redis

Response:
{
  success: true,
  booking_id: "bk_7f3a9c2e",
  bike_display_id: "Cycle-07",
  message: "Bike booked. Take pickup photo."
}
```

**Step 3: Upload Pickup Photo**
```
Student: Camera opens → "Take clear photo of bike"
App: Captures → Uploads → Server verifies

POST /api/v1/cycles/photo-pickup
{
  booking_id: "bk_7f3a9c2e",
  photo_base64: "/9j/4AAQSkZJRg..."
}

Server:
├─ Verify booking belongs to user & status='booked'
├─ Store: pickups/bk_7f3a9c2e.jpg
├─ Optional: AI check (MobileNet) → "Bike detected: 95%"
├─ AI fails → Reject: "Retake photo"
└─ Success → Generate Pickup Ticket

Response:
{
  success: true,
  ticket_qr: "<base64>",
  guard_code: "3847",
  expires_in: 300
}

App Shows:
┌─────────────────────────────────────┐
│ 🔖 PICKUP TICKET                   │
│                                    │
│ Student: Prince (2023CS0113)      │
│ Bike: Cycle-07                     │
│ Guard Code: 3 8 4 7               │
│ ⏱️ 04:59                          │
│ [QR CODE]                          │
└─────────────────────────────────────┘
```

**Step 4: Guard Verification (Guard App)**
```
Guard: Opens Guard App → [SCAN TICKET]
App: Scans → Decodes JWT → Shows:

┌─────────────────────────────────────┐
│ 🛡️ GUARD VERIFICATION              │
│                                    │
│ 👤 Prince (2023CS0113)            │
│ 🚲 Cycle-07                        │
│ 📸 Pickup Photo: [View]            │
│                                    │
│ ✅ Approve Unlock                  │
│ ❌ Reject & Flag                   │
└─────────────────────────────────────┘

Guard Actions:
├─ Checks Student ID card
├─ Views pickup photo (bike condition)
├─ Verifies bike physically there
└─ Taps ✅ Approve

POST /api/v1/guard/approve-pickup
{
  ticket_token: "...",
  guard_id: "guard_bh_01",
  action: "approve"
}

Server:
├─ Verify Guard JWT
├─ Update booking: status='picked_up', picked_up_at=NOW()
├─ Update bike: status='in_use'
├─ Clear Redis expiry
└─ Response: { success: true }

Guard: Unlocks physical lock with key
Student: Takes bike & rides
```

---

### **PHASE 2: RIDE (Student App)**

**Report Issue (Optional)**
```
Student: App → [⚠️ REPORT ISSUE]
Select: "Flat Tire" → Camera → Photo → Submit

POST /api/v1/ride/report
{
  booking_id: "bk_7f3a9c2e",
  issue_type: "flat_tire",
  photo_base64: "..."
}

Server:
├─ Flag bike: status='maintenance'
├─ Notify admin via WebSocket
└─ Response: { success: true }
```

---

### **PHASE 3: RETURN (Student App)**

**Step 1: Scan Return Stand**
```
Student: At Main Gate → [SCAN RETURN STAND QR]

POST /api/v1/cycles/return-scan
{
  encrypted_return_stand_qr: "...",
  booking_id: "bk_7f3a9c2e"
}

Server:
├─ Decrypt stand QR
├─ Verify booking status='picked_up'
├─ Check ride_duration > 2 min
└─ Response: { stand_verified: true, message: "Take return photo" }
```

**Step 2: Upload Return Photo**
```
Student: Camera → "Photo of bike at stand"

POST /api/v1/cycles/photo-return
{
  booking_id: "bk_7f3a9c2e",
  photo_base64: "..."
}

Server:
├─ Store: returns/bk_7f3a9c2e.jpg
├─ AI check: "Bike detected at stand: 92%"
├─ Update booking: return_stand='stand_mg_12'
├─ Generate Return Ticket

Response:
{
  success: true,
  ticket_qr: "<base64>",
  duration_minutes: 45,
  message: "Return ticket generated"
}

App Shows:
┌─────────────────────────────────────┐
│ ✅ RETURN TICKET                   │
│                                    │
│ Bike: Cycle-07                     │
│ Duration: 45 min                   │
│ 📸 Before/After Photos: [View]    │
│ [QR CODE]                          │
│ Show to guard for approval         │
└─────────────────────────────────────┘
```

**Step 3: Guard Return Verification (Guard App)**
```
Guard: Scans Return Ticket → Sees:

┌─────────────────────────────────────┐
│ 🛡️ RETURN VERIFICATION             │
│                                    │
│ 👤 Prince (2023CS0113)            │
│ 🚲 Cycle-07                        │
│ ⏱️ 45 min ride                     │
│                                    │
│ 📸 BEFORE: [View Photo]            │
│ 📸 AFTER:  [View Photo]            │
│                                    │
│ ✅ Approve & Collect Keys          │
│ ⚠️ Flag Damage                     │
└─────────────────────────────────────┘

Guard: Compares photos → Checks bike condition
If OK → Taps ✅ Approve

POST /api/v1/guard/approve-return
{
  ticket_token: "...",
  guard_id: "guard_mg_01",
  action: "approve",
  damage: false
}

Server:
├─ Update booking: status='returned', returned_at=NOW()
├─ Update bike: status='available', current_stand='stand_mg_12'
├─ Clear Redis: DEL user:active:2023CS0113
├─ INCR stand:available:stand_mg_12
├─ Check violations (late, fast, no-show)
└─ Response: { success: true }

Guard: Collects keys → Student leaves
```

---

## 🔧 **GUARD APP SPECIFICATION**

### **Guard Authentication**
```javascript
POST /api/v1/guard/auth
{
  guard_id: "guard_bh_01",
  password: "..." // Or OTP from admin
}

Response:
{
  guard_jwt: "eyJhbGciOiJIUzI1NiIs...",
  name: "Ramesh Kumar",
  assigned_stand: "stand_bh_01"
}
```

### **Guard Dashboard**
```
┌─────────────────────────────────────┐
│ GUARD DASHBOARD                    │
│ Ramesh - Boys Hostel               │
│                                    │
│ 🔄 Recent Scans:                   │
│ • Cycle-07 (Prince) - Approved    │
│ • Cycle-12 (Anjali) - Pending     │
│                                    │
│ [SCAN TICKET] [VIEW LOGS]         │
└─────────────────────────────────────┘
```

### **Guard Scan Ticket**
```javascript
POST /api/v1/guard/scan
{
  ticket_qr: "..."
}

Response:
{
  type: "pickup"|"return",
  student: { name, reg_no },
  bike: { display_id },
  photos: { pickup_url, return_url },
  duration: 45,
  guard_code: "3847"
}
```

### **Guard Approve/Reject**
```javascript
POST /api/v1/guard/action
{
  ticket_token: "...",
  action: "approve"|"reject",
  reason: "damage"|"no_show"|"suspicious"
}
```

---

## 📊 **ADMIN PANEL - REAL-TIME DASHBOARD**

```javascript
GET /api/v1/admin/dashboard

Response:
{
  overview: {
    total_cycles: 20,
    available: 12,
    in_use: 8,
    maintenance: 0,
    violations_today: 3
  },
  stands: [
    {
      stand_id: "stand_bh_01",
      name: "Boys Hostel",
      capacity: 10,
      available: 7,
      booked: 1,
      in_use: 2,
      last_updated: "14:36:22"
    },
    // ...
  ],
  recent_bookings: [
    {
      booking_id: "bk_7f3a9c2e",
      student: "Prince (2023CS0113)",
      bike: "Cycle-07",
      status: "picked_up",
      guard: "Ramesh",
      photo_url: "/photos/pickups/..."
    }
  ],
  violations: [
    {
      reg_no: "2023CS0113",
      type: "late_return",
      points: 1,
      time: "15:45:00"
    }
  ]
}
```

**Admin Actions:**
- Ban/Unban users
- Flag bikes for maintenance
- Override bookings (emergency)
- Download reports (CSV/Excel)
- View photo evidence

---

## 🛡️ **COMPLETE SECURITY CASES**

### **Case 1: Fake Dynamic Ticket QR**
**Attack**: Student generates fake ticket QR
**Prevention**:
- Ticket = JWT signed with `GUARD_SECRET`
- Guard app verifies signature → Rejects fakes
- Expiry: 5 min (pickup), 10 min (return)

### **Case 2: Reusing Old Ticket**
**Attack**: Screenshot old ticket, shows to guard
**Prevention**:
- Guard app checks `exp` claim → Shows "EXPIRED" in red
- Redis state: `bike:status` must match ticket type

### **Case 3: Guard Collusion**
**Attack**: Guard approves without checking
**Prevention**:
- **Photo audit**: Admin can review all approvals
- **Random audits**: Admin cross-checks guard logs vs bookings
- **Penalty**: Guard found negligent → Disciplinary action

### **Case 4: Student Never Picks Up**
**Attack**: Books, generates ticket, but doesn't go to guard
**Prevention**:
- **Booking expires in 5 min**: Redis TTL auto-releases bike
- **3 no-shows = auto-ban**: Cron job tracks

### **Case 5: Student Returns Damaged Bike**
**Attack**: Damages bike, returns, claims it was already damaged
**Prevention**:
- **Pickup photo**: Shows bike condition before ride
- **Return photo**: Shows condition after ride
- **Guard comparison**: If damage detected → Flag user
- **Pattern detection**: >2 damage reports → Auto-ban

### **Case 6: Bike Theft (Never Returns)**
**Attack**: Takes bike, never scans return
**Prevention**:
- **Overdue detection**: Cron checks >4 hours
- **Auto-alert**: Admin gets email with student details + last photo
- **Student fined**: ₹500 penalty (manual ERP integration or ban until paid)

### **Case 7: Fake Return Photo**
**Attack**: Uploads old photo from gallery
**Prevention**:
- **App blocks gallery**: Forces camera only
- **EXIF validation**: Checks photo timestamp vs server time
- **AI detection**: Ensures bike is at stand (background matching)

### **Case 8: Multiple Concurrent Bookings**
**Attack**: Tries to book multiple bikes
**Prevention**:
- **Redis check**: `user:active_booking:{reg_no}` exists → Reject
- **DB check**: `SELECT * FROM bookings WHERE reg_no=? AND status IN ('booked','picked_up')`

### **Case 9: Stand QR Tampering**
**Attack**: Replaces stand QR with fake
**Prevention**:
- **Decryption fails**: App shows "Invalid QR"
- **Physical security**: Laminated + taped + periodic checks

### **Case 10: API Spamming**
**Attack**: Bots booking all bikes
**Prevention**:
- **Rate limit**: 1 booking/minute per user (Redis)
- **Rate limit**: 10 requests/minute per IP
- **Device fingerprint**: Blocks emulator farms

---

## 💰 **FINAL COST BREAKDOWN**

| Component | Solution | Cost (₹) |
|-----------|----------|----------|
| **Backend Server** | College Computer / Railway Free | 0 |
| **Database** | PostgreSQL (College Server) | 0 |
| **Redis** | Upstash (10k cmds/day free) | 0 |
| **QR Codes** | Print 23 QRs + Laminate | 50 |
| **Photo Storage** | Local Disk (500GB) | 0 |
| **Student App** | React Native (Development) | 0 |
| **Guard App** | React Native / PWA (Lightweight) | 0 |
| **Total** | | **₹50** |

---

## 🗓️ **PRODUCTION DEPLOYMENT PLAN**

### **Week 1: Development**
- **Day 1-2**: Backend + Auth + QR Crypto
- **Day 3-4**: Booking Flow + Redis Locks
- **Day 5-6**: Photo Upload + Ticket Generation
- **Day 7**: Guard App MVP

### **Week 2: Testing**
- **Day 8-9**: Mobile App (Student) - Scan & Photo
- **Day 10-11**: Mobile App (Guard) - Scan & Approve
- **Day 12**: Integration Testing (2 bikes, 1 stand)
- **Day 13**: Load Testing (50 concurrent bookings)
- **Day 14**: Security Audit (penetration testing)

### **Week 3: Pilot**
- **Day 15**: Deploy to Boys Hostel (5 bikes)
- **Day 16-17**: Train 2 guards on app usage
- **Day 18-20**: Monitor & collect feedback
- **Day 21**: Fix bugs & optimize

### **Week 4: Full Rollout**
- **Day 22**: Generate all 20 bike QRs
- **Day 23**: Deploy to all stands
- **Day 24**: Train all guards (30 min session)
- **Day 25**: Go Live! 🚀
- **Day 26-28**: Monitor & support
- **Day 29-30**: Handover to student council

---

## 📈 **SUCCESS METRICS**

| Metric | Target | Measurement |
|--------|--------|-------------|
| **Booking Time** | < 3 min | Avg from scan to ticket |
| **Return Time** | < 2 min | Avg from scan to approval |
| **Guard Approval Time** | < 30 sec | Time to scan & approve |
| **Availability Accuracy** | 100% | Redis vs Physical count |
| **No-Show Rate** | < 5% | Expired bookings / Total |
| **Violation Rate** | < 10% | Violations / Total rides |
| **User Satisfaction** | > 4.5/5 | In-app survey |
| **Security Incidents** | 0 | Attempted hacks detected |

---

## 🎯 **TECHNICAL STACK**

| Component | Technology | Version | Rationale |
|-----------|------------|---------|-----------|
| **Backend** | Node.js + Fastify | 20.x | High perf, low latency |
| **Database** | PostgreSQL | 15.x | ACID compliance, reliable |
| **Cache** | Redis | 7.x | Atomic ops, pub/sub |
| **Auth** | Google OAuth 2.0 | - | Secure, no passwords |
| **Encryption** | AES-256-CBC | - | Industry standard |
| **JWT** | jsonwebtoken | 9.x | Battle-tested |
| **Mobile** | React Native | 0.72 | Cross-platform |
| **Guard App** | React Native / PWA | - | Lightweight, fast |
| **Storage** | Local Filesystem | - | Zero cost |
| **WebSocket** | Socket.io | 4.x | Real-time sync |
| **AI Check** | MobileNet (TF.js) | - | Free, local inference |

---

## 📱 **APP UI/UX SPECIFICATION**

### **Student App Screens**

**1. Login Screen**
```
┌─────────────────────────────────────┐
│ SVCE CYCLE ACCESS                  │
│                                    │
│ [🎓 Login with College Mail]      │
│ (2023CS0113@svce.ac.in)            │
└─────────────────────────────────────┘
```

**2. Home Screen**
```
┌─────────────────────────────────────┐
│ Welcome Prince                     │
│                                    │
│ [📍 SCAN STAND QR]                 │
│                                    │
│ 🚴‍♂️ Active: Cycle-07 (45 min)    │
│ [View Ticket] [Report Issue]      │
│                                    │
│ 📜 Ride History                    │
│ ⚙️ Settings                        │
└─────────────────────────────────────┘
```

**3. Ticket Screen (Pickup)**
```
┌─────────────────────────────────────┐
│ 🔖 PICKUP TICKET                   │
│ ⏱️ 04:59 (expires soon)          │
│                                    │
│ 👤 Prince (2023CS0113)            │
│ 🚲 Cycle-07                        │
│ 📍 Boys Hostel - Zone A            │
│                                    │
│ 📸 Pickup Photo: [View]            │
│                                    │
│ Guard Code: 3 8 4 7               │
│                                    │
│ [QR CODE - Full Screen]            │
│                                    │
│ [Cancel Booking]                   │
└─────────────────────────────────────┘
```

**4. Ticket Screen (Return)**
```
┌─────────────────────────────────────┐
│ ✅ RETURN TICKET                   │
│                                    │
│ 🚲 Cycle-07                        │
│ ⏱️ 45 min ride                     │
│ 📍 Main Gate - Slot 12             │
│                                    │
│ 📸 BEFORE: [View]                  │
│ 📸 AFTER:  [View]                  │
│                                    │
│ [QR CODE - Full Screen]            │
│ Show to guard for approval         │
└─────────────────────────────────────┘
```

### **Guard App Screens**

**1. Login Screen**
```
┌─────────────────────────────────────┐
│ GUARD LOGIN                        │
│                                    │
│ Guard ID: [guard_bh_01]            │
│ Password: [********]               │
│                                    │
│ [LOGIN]                            │
└─────────────────────────────────────┘
```

**2. Scan Ticket Screen**
```
┌─────────────────────────────────────┐
│ SCAN STUDENT TICKET                │
│                                    │
│ [CAMERA VIEWFINDER]                │
│                                    │
│ Align QR within frame              │
└─────────────────────────────────────┘
```

**3. Verification Screen (Pickup)**
```
┌─────────────────────────────────────┐
│ 🛡️ PICKUP VERIFICATION             │
│                                    │
│ 👤 Prince (2023CS0113)            │
│ 🚲 Cycle-07                        │
│ 📍 Boys Hostel - Zone A            │
│                                    │
│ 📸 Pickup Photo: [View Full]      │
│                                    │
│ Guard Code: 3 8 4 7               │
│                                    │
│ [✅ APPROVE UNLOCK]                │
│ [❌ REJECT & REPORT]               │
└─────────────────────────────────────┘
```

**4. Verification Screen (Return)**
```
┌─────────────────────────────────────┐
│ 🛡️ RETURN VERIFICATION             │
│                                    │
│ 👤 Prince (2023CS0113)            │
│ 🚲 Cycle-07                        │
│ ⏱️ 45 min                          │
│                                    │
│ 📸 BEFORE: [View]                  │
│ 📸 AFTER:  [View]                  │
│                                    │
│ [✅ APPROVE RETURN]                │
│ [⚠️ FLAG DAMAGE]                   │
└─────────────────────────────────────┘
```

---

## 🚀 **DEPLOYMENT CHECKLIST**

- [ ] **Backend**: Deploy on college server (port 443, SSL cert)
- [ ] **Database**: Run migrations, create indexes
- [ ] **Redis**: Setup instance, configure TTL policies
- [ ] **QR Codes**: Print 23 QRs, laminate, paste securely
- [ ] **Guard App**: Install on 3 guard phones, issue credentials
- [ ] **Student App**: Build APK/IPA, distribute via WhatsApp
- [ ] **Admin Panel**: Deploy on admin computer, issue admin JWT
- [ ] **Training**: 30-min session for guards (scan → verify → approve)
- [ ] **Pilot**: Test with 5 bikes for 2 days
- [ ] **Rollout**: Full deployment, monitor logs

---

## 📝 **PROJECT CHANGELOG**

### **v0.1.0** 
- **Backend**
    - Created the main directories for the backend server and mobile application.
    - Initialized a new Go module within the `server` directory.
    - Created the initial `main.go` file for the server.
    - Installed `goose` for database migrations.
    - Designed and created the initial PostgreSQL database schema (`users`, `stands`, `bikes`, `bookings`, `violations` tables) in a `goose` migration file.
    - Integrated Gin web framework.
    - Implemented configuration management, loading settings from environment variables and a `.env` file.
- **Frontend**
    - Switched mobile app development to Flutter.
    - Manually downloaded and extracted Flutter SDK.
    - Manually downloaded and extracted Android Command-line Tools.
    - Installed Android SDK components (platform-tools, build-tools, platforms).
    - Configured Flutter to use JDK 17 explicitly.
    - Initialized the Flutter project (`mobile_app`) in the `mobile-app` directory.

---

## 📞 **SUPPORT & MAINTENANCE**

- **Level 1**: Student volunteers (app usage help)
- **Level 2**: System admin (guard login issues, ban appeals)
- **Level 3**: Developer team (backend bugs, encryption issues)
- **Emergency**: Admin can manually override any booking via DB

---

## 🎓 **CONCLUSION**

This system provides **bank-level security** with **zero hardware investment** by combining:
1. **Cryptography** (AES-256) for QR integrity
2. **Distributed Locks** (Redis) for race prevention
3. **Photo Evidence** for fraud deterrence
4. **Human Verification** for final approval
5. **Automated Violations** for policy enforcement

**Result**: A system that **cannot be hacked** by hostel peers, **costs ₹50**, and **scales to 100+ bikes** with zero changes.

---

## 💻 **READY TO BUILD?**

**Next Step**: I'll provide the **complete GitHub repository template** with:
- Backend boilerplate (Fastify + JWT + Redis)
- Student App starter (React Native + Camera)
- Guard App starter (React Native + QR Scanner)
- QR generation script
- Database migration files

**Say "Da Boi, repo kudu!" and let's code!** 🚀💻
