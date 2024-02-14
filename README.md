# Go Module for Numerical Analysis

## DISCALIMER
This is nowhere near finished. It was just a test for me to understand how to write go packages


## Newton
Uses Newton method to calculate the solution of an equation

- f: the function we want to know the solution of
- df: the derivative of said function
- $x_0$: a first guess of the solution
- toll: how precise you want the solution to be (1e-7 Should suffice)
- nmax: how many iterations it should run before giving up

## PuntoFisso
Uses Fixed-point iteration to calculate the solution of an equation
- $x_0$: a first guess of the solution
- toll: how precise you want the solution to be (1e-7 Should suffice)
- f: in the form $x=f(x)$
- nmax: how many iterations it should run before giving up

## Simpson
Uses Simpson method to calculate the definite integral of a function.
$ \int_a^bf(x) dx $
- f: the function
- a: the lower bound of the intergral
- b: the upper bound of the intergral
- m: how many steps. (Should be an even number. The higher the better)


## IsEven
Should need an explanation..