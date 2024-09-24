use crate::binary::encoder::encode_value;
use crate::buffer::ReverseWriterFactory;
use crate::codec::Codec;
use crate::decoder::DecodeError;
use crate::encoder::EncodeError;
use crate::mem::MemoryManager;
use crate::value::Value;

mod encoder;
mod decoder;

pub struct NativeBinaryCodec;

impl Codec for NativeBinaryCodec {
    fn encode_value<'a, V: Value<'a>, F: ReverseWriterFactory>(value: &V, writer_factory: &F) -> Result<F::Writer::Output, EncodeError> {
        encode_value(value, writer_factory)
    }

    fn decode_value<'b, 'a: 'b, V: Value<'a>>(input: &'a [u8], memory_manager: &'b MemoryManager<'a, 'a>) -> Result<V, DecodeError> {
        decoder::decode_value(input, memory_manager)
    }
}