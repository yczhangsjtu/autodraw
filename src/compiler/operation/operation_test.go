package operation

import "testing"

func BenchmarkOperationPrint(b *testing.B) {
	OperationPrint(NewOperation(LINE));
	OperationPrint(NewOperation(RECT));
	OperationPrint(NewOperation(CIRCLE));
	OperationPrint(NewOperation(OVAL));
}
