import 'package:flutter/material.dart';

class HomeScreen extends StatefulWidget {
  final String role;
  const HomeScreen({super.key, required this.role});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  String? _imagePath;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Welcome, ${widget.role}!'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('This is the ${widget.role} dashboard.'),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: () async {
                final imagePath = await Navigator.of(context).pushNamed('/camera');
                if (imagePath is String) {
                  setState(() {
                    _imagePath = imagePath;
                  });
                }
              },
              child: const Text('Take Picture'),
            ),
            if (_imagePath != null) ...[
              const SizedBox(height: 20),
              Text('Image Path: $_imagePath'),
            ],
          ],
        ),
      ),
    );
  }
}
