package asyncgroup_sync

func RunTxn[D any](base DaxBase, logics ...func(dax D) Err) Err {
	base.begin()

	dax := base.(D)
	err := Ok()

	for _, logic := range logics {
		err = logic(dax)
		if err.IsNotOk() {
			break
		}
	}

	if err.IsOk() {
		err = base.commit()
	}

	if !err.IsOk() {
		base.rollback()
	}

	base.end()

	return err
}

func Txn[D any](base DaxBase, logics ...func(dax D) Err) func() Err {
	return func() Err {
		return RunTxn[D](base, logics...)
	}
}
