import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:princetheprogrammerbtw/TheBiCycleApp/mobile-app/mobile_app/lib/screens/login_screen.dart';
import 'package:princetheprogrammerbtw/TheBiCycleApp/mobile-app/mobile_app/lib/screens/role_gate.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'SVCE Cycle Access',
      theme: _buildThemeData(),
      initialRoute: '/',
      routes: {
        '/': (context) => const LoginScreen(),
        '/role_gate': (context) => const RoleGate(),
      },
      debugShowCheckedModeBanner: false,
    );
  }

  ThemeData _buildThemeData() {
    const primaryColor = Color(0xFF6200EE);
    const secondaryColor = Color(0xFF03DAC6);

    return ThemeData(
      useMaterial3: true,
      colorScheme: ColorScheme.fromSeed(
        seedColor: primaryColor,
        primary: primaryColor,
        secondary: secondaryColor,
        background: const Color(0xFFF4F6F8),
        surface: Colors.white,
        onPrimary: Colors.white,
        onSecondary: Colors.black,
        onBackground: Colors.black,
        onSurface: Colors.black,
        error: const Color(0xFFB00020),
        onError: Colors.white,
      ),
      textTheme: GoogleFonts.poppinsTextTheme(),
      appBarTheme: const AppBarTheme(
        backgroundColor: primaryColor,
        foregroundColor: Colors.white,
        elevation: 4,
      ),
      elevatedButtonTheme: ElevatedButtonThemeData(
        style: ElevatedButton.styleFrom(
          backgroundColor: primaryColor,
          foregroundColor: Colors.white,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(8),
          ),
          padding: const EdgeInsets.symmetric(vertical: 16),
        ),
      ),
    );
  }
}