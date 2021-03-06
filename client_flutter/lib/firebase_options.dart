// File generated by FlutterFire CLI.
// ignore_for_file: lines_longer_than_80_chars
import 'package:firebase_core/firebase_core.dart' show FirebaseOptions;
import 'package:flutter/foundation.dart'
    show defaultTargetPlatform, kIsWeb, TargetPlatform;

/// Default [FirebaseOptions] for use with your Firebase apps.
///
/// Example:
/// ```dart
/// import 'firebase_options.dart';
/// // ...
/// await Firebase.initializeApp(
///   options: DefaultFirebaseOptions.currentPlatform,
/// );
/// ```
class DefaultFirebaseOptions {
  static FirebaseOptions get currentPlatform {
    if (kIsWeb) {
      return web;
    }
    // ignore: missing_enum_constant_in_switch
    switch (defaultTargetPlatform) {
      case TargetPlatform.android:
        return android;
      case TargetPlatform.iOS:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for ios - '
          'you can reconfigure this by running the FlutterFire CLI again.',
        );
      case TargetPlatform.macOS:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for macos - '
          'you can reconfigure this by running the FlutterFire CLI again.',
        );
    }

    throw UnsupportedError(
      'DefaultFirebaseOptions are not supported for this platform.',
    );
  }

  static const FirebaseOptions web = FirebaseOptions(
    apiKey: 'AIzaSyDRPGLNdoyC-FPMHkuUQcO4cCfO77tHEpE',
    appId: '1:311242090009:web:21c4244271afa337076c88',
    messagingSenderId: '311242090009',
    projectId: 'koku-e0b26',
    authDomain: 'koku-e0b26.firebaseapp.com',
    storageBucket: 'koku-e0b26.appspot.com',
    measurementId: 'G-VE73JZ0ST8',
  );

  static const FirebaseOptions android = FirebaseOptions(
    apiKey: 'AIzaSyAgFsSycDvWDFuQ_WR8tRDJugUzMpmQgU0',
    appId: '1:311242090009:android:ae30c254876727be076c88',
    messagingSenderId: '311242090009',
    projectId: 'koku-e0b26',
    storageBucket: 'koku-e0b26.appspot.com',
  );
}
