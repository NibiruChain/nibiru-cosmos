package types

// ----------------------------------------------------------------------------
// New EVM
// ----------------------------------------------------------------------------

// keys should be unexported, unique types
type lastErrApplyEvmMsg struct{}

// Get/Set helpers. Value receiver is fine: it mutates the pointed slot.
func (c Context) LastErrApplyEvmMsg() error {
	if c.lastErrApplyEvmMsg == nil {
		return nil
	}
	return c.lastErrApplyEvmMsg.evmErr
}
func (c Context) WithLastErrApplyEvmMsg(e error) {
	if c.lastErrApplyEvmMsg == nil {
		return
	}
	c.lastErrApplyEvmMsg.evmErr = e
}

func (c Context) IsEvmTx() bool { return c.isEvmTx }
func (c Context) WithIsEvmTx(isEvmTx bool) Context {
	c.isEvmTx = isEvmTx
	return c
}
