import 'package:flutter/material.dart';
import 'package:mobile_app/screens/home_screen.dart';

class RoleGate extends StatelessWidget {
  const RoleGate({super.key});

  @override
  Widget build(BuildContext context) {
    // TODO: Implement actual role checking logic (e.g., from an auth provider)
    const userRole = 'student'; // Mock role

    // For now, we only have one home screen, but this can be expanded.
    switch (userRole) {
      case 'student':
        return const HomeScreen(role: userRole);
      case 'guard':
        // return const GuardHomeScreen();
        return const HomeScreen(role: userRole); // Placeholder
      default:
        // Handle unknown role, maybe navigate back to login
        return const Scaffold(
          body: Center(
            child: Text('Unknown user role. Please log in again.'),
          ),
        );
    }
  }
}
