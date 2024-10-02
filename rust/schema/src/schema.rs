use crate::message::MessageDescriptor;
use crate::oneof::OneOfType;
use crate::r#enum::EnumType;
use crate::state_object::StateObjectType;
use crate::structs::StructType;

#[non_exhaustive]
#[derive(Debug, Clone, Eq, PartialEq)]
pub enum SchemaType<'a> {
    Struct(StructType<'a>),
    Enum(EnumType<'a>),
    OneOf(OneOfType<'a>),
    StateObjectType(StateObjectType<'a>),
}

impl <'a> SchemaType<'a> {
    pub const fn name(&self) -> &'a str {
        match self {
            SchemaType::Struct(s) => s.name,
            SchemaType::Enum(e) => e.name,
            SchemaType::OneOf(o) => o.name,
            SchemaType::StateObjectType(s) => s.name,
        }
    }
}

impl <'a> PartialOrd for SchemaType<'a> {
    fn partial_cmp(&self, other: &Self) -> Option<core::cmp::Ordering> {
        self.name().partial_cmp(other.name())
    }
}

impl <'a> Ord for SchemaType<'a> {
    fn cmp(&self, other: &Self) -> core::cmp::Ordering {
        self.name().cmp(other.name())
    }
}

#[non_exhaustive]
#[derive(Debug, Clone, Eq, PartialEq, Default)]
pub struct Schema<'a> {
    types: &'a [SchemaType<'a>],
    messages: &'a [MessageDescriptor<'a>],
}

impl Schema<'static> {
    pub const fn add(&self, schema_type: SchemaType<'static>) -> Self {
        todo!()
    }
}

pub trait HasSchema {
    const SCHEMA: Schema<'static>;
}