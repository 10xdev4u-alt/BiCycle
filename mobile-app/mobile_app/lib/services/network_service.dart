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

  // TODO: Add other methods for other endpoints
}
