import 'package:app/manager.dart';
import 'package:app/screen.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_telegram_miniapp/flutter_telegram_miniapp.dart';

void main() {
  runApp(
    MultiBlocProvider(
      providers: [BlocProvider<BlocBloc>(create: (context) => BlocBloc()..add(Increace()))],
      child: const MyApp(),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Telegram Mini App Example',
      theme: ThemeData(primarySwatch: Colors.blue),
      home: MainScreen(),
    );
  }
}
