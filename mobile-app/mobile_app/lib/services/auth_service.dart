import 'package:google_sign_in/google_sign_in.dart';
import 'package:url_launcher/url_launcher.dart';

class AuthService {
  final String _baseUrl = 'http://localhost:8080'; // TODO: Make configurable

  Future<void> signInWithGoogle() async {
    final url = Uri.parse('$_baseUrl/auth/google');
    if (await canLaunchUrl(url)) {
      await launchUrl(url, mode: LaunchMode.externalApplication);
    } else {
      throw 'Could not launch $url';
    }
  }

  Future<void> signOut() async {
    // TODO: Implement actual sign out logic, including clearing local JWT
    print("User signed out.");
  }
}
