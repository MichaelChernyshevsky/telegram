import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

class BlocState extends Equatable {
  final int counter;

  const BlocState({required this.counter});

  @override
  List<Object?> get props => [counter];

  BlocState copyWith({int? counter}) {
    return BlocState(counter: counter ?? this.counter);
  }

  factory BlocState.initial() {
    return BlocState(counter: 11);
  }
}

abstract class BlocEvent extends Equatable {}

class Increace extends BlocEvent {
  @override
  List<Null> get props => [];
}

class BlocBloc extends Bloc<BlocEvent, BlocState> {
  BlocBloc() : super(BlocState.initial()) {
    on<Increace>(_increace);
  }
  Future<void> _increace(Increace event, Emitter<BlocState> emit) async {
    emit(state.copyWith(counter: state.counter + 1));
  }
}
