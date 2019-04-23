# Protist (alpha software)

<img src="https://github.com/ganehag/protist/raw/master/docs/images/logo.png" width="200">

---

*Computation of values using a chain of logic and math.*

Protist is a simple, yet flexible engine to transform values using a chain of functions.

In some ways it behaves similar to Functional Programming (FP). It does;

- Avoid shared state
- Avoid mutating state
- Avoid side effects

Protist requires you to solve problems using existing building blocks. As it;

- Only supports pure functions
- And it doesn't allow loops (no `for` or `while`)

---

## Chains

A chain is a series of functions where the `output` from one or serveral functions acts as `input` to another
function, or multiple functions. *Each function is executed in order and this order is important.*


**NOTE: Protist doesn't currently support input in the following format. This is just an example.**

```
1: const 10   // constant value 10
2: const 20   // constant value 20
3: add $1 $2  // add (+) line 1 and line 2 -> 30
4: sub $1 $3  // subtract (-) line 3 from line 1 -> -20
5: add $1 $2  // again, add (+) line 1 and line 2 -> 30
```

The result is always tracked by the instance of the function. It is therefore impossible to overwrite a result of a function.


## Backend

While Protist currently doesn't have a well defined backend for chain declaration. It does support a very basic JSON structure
and also provides a way for backend plug-ins to be developed.


## GUI

Plans for a GUI exists and it will probably look something similar to this mock-up.

![Gui Mock-up][gui-mockup]







[gui-mockup]: docs/images/gui_mockup.png
