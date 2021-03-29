# Concurrency » Future

## Description

This design pattern describe an object that acts as a proxy for a result that is
initially unknown, usually because the computation of its value is yet incomplete.

The following is a solution with row code. First function in case of failure and
second function in case of success. Here, for semplicity is just checked that the
number is equal or not to 42. In a real world example there could be an http call
or a query to database. Let's imagine whatever you want.

```go
var wg sync.WaitGroup
wg.Add(1)

go func(num int) {
  if num != 42 {
    // in case of failure
    func() {
      fmt.Println("wrong answer")
      wg.Done()
    }()
  } else {
    // in case of success
    func() {
      fmt.Println("right answer")
      wg.Done()
    }()
  }
  fmt.Println(num)
}(42)

wg.Wait()
```

## Implementation

Now let's build a real implementation of the pattern. First of all we
should define some functions for the future:

 - SuccessFunc
 - FailureFunc
 - ExecuteFunc

```go
type SuccessFunc func(string)
type FailureFunc func(error)
type ExecuteFunc func(int) (string, error)
```

In this particular example, the execute function receive an integer. In this
pattern, this function, should always return a string and an error. In case of
success, string is passed to SuccessFunc. On the other case, error will be sent
to FailureFunc.

The following code represent the Subject in the pattern. In a real world
example it could be better to chose a more semantic name.

```go
type Subject struct {
	success SuccessFunc
	failure FailureFunc
}

func (s *Subject) Success(f SuccessFunc) *Subject {
	s.success = f
	return s
}

func (s *Subject) Failure(f FailureFunc) *Subject {
	s.failure = f
	return s
}
```

The key point of the pattern is here. Like in barrier pattern is used a
WaitGroup. Just for this example, a random number is generated. This should be
replaced with a real domain algorithm.

The fact is that `ExecuteFunc` will be executed and it will return a message or
an error.

```go
func (s *Subject) Execute(f ExecuteFunc) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func(s *Subject) {
		r := rand.NewSource(time.Now().UnixNano())
		n := rand.New(r)

		str, err := f(n.Intn(200))

		if err != nil {
			s.failure(err)
			wg.Done()
		} else {
			s.success(str)
			wg.Done()
		}
	}(s)

	wg.Wait()
}
```

And finally, the main function where our promise is a struct called Subject.
This struct implement Success, Failure and Execute functions.

```go
func main() {
	s := Subject{}

	s.Success(func(m string) {
		fmt.Println("SUCCESS: ", m)
	}).Failure(func(e error) {
		fmt.Println("FAILURE: ", e)
	}).Execute(func(num int) (string, error) {
		if num < 100 {
			return "", errors.New("too low number")
		} else {
			return "valid number", nil
		}
	})

}
```