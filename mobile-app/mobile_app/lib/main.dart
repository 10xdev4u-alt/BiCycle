import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:princetheprogrammerbtw/TheBiCycleApp/mobile-app/mobile_app/lib/screens/login_screen.dart';
import 'package:princetheprogrammerbtw/TheBiCycleApp/mobile-app/mobile_app/lib/screens/role_gate.dart';
import 'packagepackage:princetheprogrammerbtw/TheBiCycleApp/mobile-app/mobile_app/lib/screens/camera_screen.dart';
import 'packagepackage:princetheprogrammerbtw/TheBiCycleApp/mobile-app/mobile_app/lib/screens/scan_qr_screen.dart';

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
        '/camera': (context) => const CameraScreen(),
        '/scan_qr': (context) => const ScanQrScreen(),
      },
      debugShowCheckedModeBanner: false,
    );
  }

    ThemeData _buildThemeData() {

      const primaryColor = Color(0xFF003B46);

      const secondaryColor = Color(0xFF07575B);

      const accentColor = Color(0xFF66A5AD);

      const backgroundColor = Color(0xFFF4F6F8);

      const surfaceColor = Colors.white;

      const onPrimaryColor = Colors.white;

      const onBackgroundColor = Color(0xFF121212);

  

      final baseTheme = ThemeData.from(

        colorScheme: ColorScheme(

          primary: primaryColor,

          secondary: secondaryColor,

          surface: surfaceColor,

          background: backgroundColor,

          error: const Color(0xFFB00020),

          onPrimary: onPrimaryColor,

          onSecondary: onPrimaryColor,

          onSurface: onBackgroundColor,

          onBackground: onBackgroundColor,

          onError: Colors.white,

          brightness: Brightness.light,

        ),

        useMaterial3: true,

      );

  

      return baseTheme.copyWith(

        textTheme: GoogleFonts.poppinsTextTheme(baseTheme.textTheme),

        appBarTheme: const AppBarTheme(

          backgroundColor: primaryColor,

          foregroundColor: onPrimaryColor,

          elevation: 0,

          centerTitle: true,

        ),

        elevatedButtonTheme: ElevatedButtonThemeData(

          style: ElevatedButton.styleFrom(

            backgroundColor: accentColor,

            foregroundColor: onPrimaryColor,

            shape: RoundedRectangleBorder(

              borderRadius: BorderRadius.circular(30),

            ),

            padding: const EdgeInsets.symmetric(vertical: 16, horizontal: 32),

          ),

        ),

        floatingActionButtonTheme: const FloatingActionButtonThemeData(

          backgroundColor: accentColor,

          foregroundColor: onPrimaryColor,

        ),

      );

    }

  }

  