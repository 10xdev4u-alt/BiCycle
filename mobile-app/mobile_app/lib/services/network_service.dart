import 'dart:convert';
import 'package:http/http.dart' as http;

class NetworkService {
  final String _baseUrl = 'http://localhost:8080/api/v1';

  Future<Map<String, dynamic>> scanStand(String encryptedQr) async {
    final response = await http.post(
      Uri.parse('$_baseUrl/stand/scan'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, String>{
        'encrypted_stand_qr': encryptedQr,
      }),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to scan stand.');
    }
  }

  Future<Map<String, dynamic>> scanAndBookCycle(
      String encryptedQr, String standId) async {
    final response = await http.post(
      Uri.parse('$_baseUrl/cycles/scan-book'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, String>{
        'encrypted_bike_qr': encryptedQr,
        'current_stand': standId,
      }),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to book cycle.');
    }
  }

  Future<Map<String, dynamic>> uploadPickupPhoto(
      String bookingId, String photoBase64) async {
    final response = await http.post(
      Uri.parse('$_baseUrl/cycles/photo-pickup'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, String>{
        'booking_id': bookingId,
        'photo_base64': photoBase64,
      }),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to upload photo.');
    }
  }

  Future<Map<String, dynamic>> scanTicket(String ticketQr) async {
    final response = await http.post(
      Uri.parse('$_baseUrl/guard/scan'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, String>{
        'ticket_qr': ticketQr,
      }),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to scan ticket.');
    }
  }

  Future<Map<String, dynamic>> approvePickup(
      String ticketToken, String guardId, String action) async {
    final response = await http.post(
      Uri.parse('$_baseUrl/guard/approve-pickup'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, String>{
        'ticket_token': ticketToken,
        'guard_id': guardId,
        'action': action,
      }),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to approve pickup.');
    }
  }

  Future<Map<String, dynamic>> approveReturn(
      String ticketToken, String guardId, String action, bool damage) async {
    final response = await http.post(
      Uri.parse('$_baseUrl/guard/approve-return'),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, dynamic>{
        'ticket_token': ticketToken,
        'guard_id': guardId,
        'action': action,
        'damage': damage,
      }),
    );

    if (response.statusCode == 200) {
      return jsonDecode(response.body);
    } else {
      throw Exception('Failed to approve return.');
    }
  }
}
