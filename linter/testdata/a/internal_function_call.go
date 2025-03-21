package a

import (
	"github.com/tbeati/stacked"

	"testdata/b"
)

func functionCallAssignmentInternal() error {
	var err error

	err = b.SingleReturn()
	if err != nil {
		return err
	}

	err = stacked.Wrap(b.SingleReturn())
	if err != nil {
		return err
	}

	err = b.SingleReturn()
	err = stacked.Wrap(err)
	if err != nil {
		return err
	}

	_, err = 0, b.SingleReturn()
	if err != nil {
		return err
	}

	_, err = 0, stacked.Wrap(b.SingleReturn())
	if err != nil {
		return err
	}

	_, err = 0, b.SingleReturn()
	err = stacked.Wrap(err)
	if err != nil {
		return err
	}

	_, err = b.MultipleReturn()
	if err != nil {
		return err
	}

	_, err = b.MultipleReturn()
	err = stacked.Wrap(err)
	if err != nil {
		return err
	}

	return nil
}

func functionCallDeclarationInternal() error {
	{
		var err = b.SingleReturn()
		if err != nil {
			return err
		}
	}

	{
		var err = stacked.Wrap(b.SingleReturn())
		if err != nil {
			return err
		}
	}

	{
		var err = b.SingleReturn()
		err = stacked.Wrap(err)
		if err != nil {
			return err
		}
	}

	{
		var _, err = 0, b.SingleReturn()
		if err != nil {
			return err
		}
	}

	{
		var _, err = 0, stacked.Wrap(b.SingleReturn())
		if err != nil {
			return err
		}
	}

	{
		var _, err = 0, b.SingleReturn()
		err = stacked.Wrap(err)
		if err != nil {
			return err
		}
	}

	{
		var _, err = b.MultipleReturn()
		if err != nil {
			return err
		}
	}

	{
		var _, err = b.MultipleReturn()
		err = stacked.Wrap(err)
		if err != nil {
			return err
		}
	}

	return nil
}

func functionCallShortDeclarationInternal() error {
	{
		err := b.SingleReturn()
		if err != nil {
			return err
		}
	}

	{
		err := stacked.Wrap(b.SingleReturn())
		if err != nil {
			return err
		}
	}

	{
		err := b.SingleReturn()
		err = stacked.Wrap(err)
		if err != nil {
			return err
		}
	}

	{
		_, err := 0, b.SingleReturn()
		if err != nil {
			return err
		}
	}

	{
		_, err := 0, stacked.Wrap(b.SingleReturn())
		if err != nil {
			return err
		}
	}

	{
		_, err := 0, b.SingleReturn()
		err = stacked.Wrap(err)
		if err != nil {
			return err
		}
	}

	{
		_, err := b.MultipleReturn()
		if err != nil {
			return err
		}
	}

	{
		_, err := b.MultipleReturn()
		err = stacked.Wrap(err)
		if err != nil {
			return err
		}
	}

	return nil
}

func functionCallReturnSingleInternal() error {
	return b.SingleReturn()
	return stacked.Wrap(b.SingleReturn())
}

func functionCallReturnMultipleInternal() (int, error) {
	return 0, b.SingleReturn()
	return 0, stacked.Wrap(b.SingleReturn())
	return b.MultipleReturn()
}

func functionCallArgumentInternal() {
	functionWithIntErrorArgument(0, b.SingleReturn())
	functionWithIntErrorArgument(0, stacked.Wrap(b.SingleReturn()))
	functionWithIntErrorArgument(b.MultipleReturn())
}

func functionCallCompositeLiteralInternal() {
	_ = structWithErrorField{
		err: b.SingleReturn(),
	}
	_ = structWithErrorField{
		err: stacked.Wrap(b.SingleReturn()),
	}

	_ = []error{b.SingleReturn()}
	_ = []error{stacked.Wrap(b.SingleReturn())}

	_ = map[string]error{"": b.SingleReturn()}
	_ = map[string]error{"": stacked.Wrap(b.SingleReturn())}
}
