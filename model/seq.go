package model

type Seq[T any] struct {
	Data T
}

type OptionalSequence[T any] struct {
	Data T
}

//func NewSeq[T any]() Seq[T] {
//	return Seq[T]{}
//}

func Filter[T any](input Seq[T], cb func(v T) bool) Seq[T] {
	return input
}

func Map[TInput any, TOutput any](input Seq[TInput], cb func(input TInput) TOutput) Seq[TOutput] {
	data := cb(input.Data)
	return Seq[TOutput]{Data: data}
}

func InnerRelationship[TFrom any, TTo any](to TTo, cb func(from TFrom, to TTo) bool) TTo {
	return to
}

func OptionalRelationship[TFrom any, TTo any](cb func(from TFrom, to TTo) bool) OptionalSequence[TTo] {
	return OptionalSequence[TTo]{}
}

func Query[T any](data T) Seq[T] {
	return Seq[T]{}
}
