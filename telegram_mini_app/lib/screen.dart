import 'package:app/manager.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class MainScreen extends StatefulWidget {
  const MainScreen({super.key});

  @override
  State<MainScreen> createState() => _MainScreenState();
}

class _MainScreenState extends State<MainScreen> {
  late BlocBloc _bloc;

  @override
  void initState() {
    super.initState();
    _bloc = BlocProvider.of<BlocBloc>(context);
  }

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<BlocBloc, BlocState>(
      builder: (context, state) {
        return Scaffold(
          body: Center(
            child: Column(
              children: [
                GestureDetector(onTap: () => _bloc.add(Increace()), child: Text('add')),
                Text(state.counter.toString()),
              ],
            ),
          ),
        );
      },
    );
  }
}
