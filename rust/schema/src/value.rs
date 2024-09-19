//! This module contains traits that must be implemented by types that can be used in the schema.

use alloc::borrow::ToOwned;
use crate::types::*;

/// Any type used directly as a message function argument or struct field must implement this trait.
/// Unlike [`Value`] it takes a lifetime parameter so value may already be borrowed where it is
/// declared.
pub trait MaybeBorrowed<'a>
where
    Self: 'a,
{
    /// The type of the value.
    type Type: Type;

    /// A dummy method for building macros until we have an actual implementation.
    fn dummy() {}
}

/// This trait describes value types that are to be used as generic parameters
/// where there is no lifetime parameter available.
/// Any types implementing this trait relate themselves to a type implementing [`MaybeBorrowed`]
/// so that generic types taking a `Value` type parameter can use a borrowed value if possible.
pub trait Value {
    /// The possibly borrowable value type this type is related to.
    type MaybeBorrowed<'a>: MaybeBorrowed<'a>;
}

impl<'a> MaybeBorrowed<'a> for u8 {
    type Type = U8T;
}
impl<'a> MaybeBorrowed<'a> for u16 {
    type Type = U16T;
}
impl<'a> MaybeBorrowed<'a> for u32 {
    type Type = U32T;
}
impl<'a> MaybeBorrowed<'a> for u64 {
    type Type = U64T;
}
impl<'a> MaybeBorrowed<'a> for u128 {
    type Type = UIntNT<16>;
}
impl<'a> MaybeBorrowed<'a> for i8 {
    type Type = I8T;
}
impl<'a> MaybeBorrowed<'a> for i16 {
    type Type = I16T;
}
impl<'a> MaybeBorrowed<'a> for i32 {
    type Type = I32T;
}
impl<'a> MaybeBorrowed<'a> for i64 {
    type Type = I64T;
}
impl<'a> MaybeBorrowed<'a> for i128 {
    type Type = IntNT<16>;
}
impl<'a> MaybeBorrowed<'a> for bool {
    type Type = Bool;
}
impl<'a> MaybeBorrowed<'a> for &'a str {
    type Type = StrT;
}

#[cfg(feature = "std")]
impl<'a> MaybeBorrowed<'a> for alloc::string::String {
    type Type = StrT;
}

impl<'a> MaybeBorrowed<'a> for simple_time::Time {
    type Type = TimeT;
}
impl<'a> MaybeBorrowed<'a> for simple_time::Duration {
    type Type = DurationT;
}
impl<'a, V: MaybeBorrowed<'a>> MaybeBorrowed<'a> for Option<V> {
    type Type = NullableT<V::Type>;
}
impl<'a, V: MaybeBorrowed<'a>> MaybeBorrowed<'a> for &'a [V]
where
    V::Type: ListElementType,
{
    type Type = ListT<V::Type>;
}

#[cfg(feature = "std")]
impl<'a, V: MaybeBorrowed<'a>> MaybeBorrowed<'a> for alloc::vec::Vec<V>
where
    V::Type: ListElementType,
{
    type Type = ListT<V::Type>;
}

#[cfg(feature = "address")]
impl<'a> MaybeBorrowed<'a> for interchain_message_api::Address {
    type Type = AddressT;
}

#[cfg(feature = "arrayvec")]
impl<'a, T: Type, V: MaybeBorrowed<'a, T>, const N: usize> MaybeBorrowed<'a, ListT<T>> for arrayvec::ArrayVec<T, N> {}
#[cfg(feature = "arrayvec")]
impl<'a, const N: usize> MaybeBorrowed<'a, StrT> for arrayvec::ArrayString<T, N> {}

impl Value for u8 {
    type MaybeBorrowed<'a> = u8;
}
impl Value for u16 {
    type MaybeBorrowed<'a> = u16;
}
impl Value for u32 {
    type MaybeBorrowed<'a> = u32;
}
impl Value for u64 {
    type MaybeBorrowed<'a> = u64;
}
impl Value for u128 {
    type MaybeBorrowed<'a> = u128;
}
impl Value for i8 {
    type MaybeBorrowed<'a> = i8;
}
impl Value for i16 {
    type MaybeBorrowed<'a> = i16;
}
impl Value for i32 {
    type MaybeBorrowed<'a> = i32;
}
impl Value for i64 {
    type MaybeBorrowed<'a> = i64;
}
impl Value for i128 {
    type MaybeBorrowed<'a> = i128;
}
impl Value for bool {
    type MaybeBorrowed<'a> = bool;
}
impl Value for str {
    type MaybeBorrowed<'a> = &'a str;
}
impl Value for simple_time::Time {
    type MaybeBorrowed<'a> = simple_time::Time;
}
impl Value for simple_time::Duration {
    type MaybeBorrowed<'a> = simple_time::Duration;
}
impl Value for interchain_message_api::Address {
    type MaybeBorrowed<'a> = interchain_message_api::Address;
}
impl<V: Value> Value for Option<V> {
    type MaybeBorrowed<'a> = Option<V::MaybeBorrowed<'a>>;
}
impl<V: Value> Value for [V]
where
        for<'a> <<V as Value>::MaybeBorrowed<'a> as MaybeBorrowed<'a>>::Type: ListElementType,
{
    type MaybeBorrowed<'a> = &'a [V::MaybeBorrowed<'a>];
}

/// ResponseValue is a trait that must be implemented by types that can be used as the return value.
pub trait ResponseValue {
    /// The type that might be borrowed.
    #[cfg(feature = "std")]
    type MaybeBorrowed<'a>: ToOwned;
    #[cfg(not(feature = "std"))]
    type MaybeBorrowed<'a>;
}
impl ResponseValue for () {
    type MaybeBorrowed<'a> = ();
}
impl<V: Value> ResponseValue for V
where
        for<'a> V::MaybeBorrowed<'a>: ToOwned,
{
    type MaybeBorrowed<'a> = V::MaybeBorrowed<'a>;
}
