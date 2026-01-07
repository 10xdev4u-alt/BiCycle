import 'package:flutter/material.dart';

class HomeScreen extends StatelessWidget {
  final String role;
  const HomeScreen({super.key, required this.role});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Welcome, $role!'),
      ),
      body: Center(
        child: Text('This is the $role dashboard.'),
      ),
    );
  }
}
