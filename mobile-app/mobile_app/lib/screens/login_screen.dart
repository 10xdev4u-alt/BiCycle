import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:princetheprogrammerbtw/TheBiCycleApp/mobile-app/mobile_app/lib/services/auth_service.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final AuthService _authService = AuthService();
  bool _isLoading = false;

  void _signIn() async {
    setState(() {
      _isLoading = true;
    });

    final user = await _authService.signInWithGoogle();

    if (mounted) {
      setState(() {
        _isLoading = false;
      });

      if (user != null) {
        Navigator.of(context).pushReplacementNamed('/role_gate');
      } else {
        // Handle sign-in failure
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Google Sign-In failed. Please try again.'),
            backgroundColor: Colors.red,
          ),
        );
      }
    }
  }

    @override

    Widget build(BuildContext context) {

      final theme = Theme.of(context);

      final isDarkMode = theme.brightness == Brightness.dark;

  

      return Scaffold(

        backgroundColor: theme.colorScheme.background,

        body: SafeArea(

          child: Center(

            child: SingleChildScrollView(

              padding: const EdgeInsets.all(24.0),

              child: Column(

                mainAxisAlignment: MainAxisAlignment.center,

                children: [

                  Icon(

                    Icons.directions_bike,

                    size: 80,

                    color: theme.colorScheme.primary,

                  ),

                  const SizedBox(height: 24),

                  Text(

                    'SVCE Cycle Access',

                    textAlign: TextAlign.center,

                    style: theme.textTheme.headlineLarge?.copyWith(

                      fontWeight: FontWeight.bold,

                      color: theme.colorScheme.primary,

                    ),

                  ),

                  const SizedBox(height: 8),

                  Text(

                    'Your campus ride, one scan away.',

                    textAlign: TextAlign.center,

                    style: theme.textTheme.titleMedium?.copyWith(

                      color: theme.colorScheme.onBackground.withOpacity(0.7),

                    ),

                  ),

                  const SizedBox(height: 48),

                  if (_isLoading)

                    const CircularProgressIndicator()

                  else

                    ElevatedButton.icon(

                      onPressed: _signIn,

                      style: theme.elevatedButtonTheme.style?.copyWith(

                        backgroundColor: MaterialStateProperty.all(theme.colorScheme.primary),

                        foregroundColor: MaterialStateProperty.all(theme.colorScheme.onPrimary),

                      ),

                      icon: const Icon(Icons.school, size: 24),

                      label: const Text(

                        'Login with College Mail',

                      ),

                    ),

                ],

              ),

            ),

          ),

        ),

      );

    }

  }

  