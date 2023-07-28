## Out of order events

Your task is to get all the tests passing

        ❯ go test -v .
        === RUN   TestHandleEventDelivery
        === RUN   TestHandleEventDelivery/in_order
        === RUN   TestHandleEventDelivery/out_of_order
        sentence_test.go:30: Wanted "the quick brown fox" but got " fox brown quickthe"
        --- FAIL: TestHandleEventDelivery (0.00s)
        --- PASS: TestHandleEventDelivery/in_order (0.00s)
        --- FAIL: TestHandleEventDelivery/out_of_order (0.00s)
        FAIL
        FAIL	adsd/outoforder	0.106s
        FAIL

You'll probably spend most time in `sentence.go`, `evt.TextAppended` and `evt.Event`