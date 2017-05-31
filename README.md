# Algorithms in Golang

A lot of introductory algorithms courses use a dynamic language like Python
or Javascript because there is an idea that it will be "easier" for people to
get going with. While there may be some truth to that, I believe that it is
actually harder to learn algorithms with a dynamic language because too many
important details are hidden, obscured or inaccessible. Things that should be
fast are sometimes slow, things that should be simple are sometimes complex
and trade-offs that are subtle are decided for you.

On the other hand, trying to investigate algorithms in low-level static
language like C presents too many tooling challenges, language quirks and
system setup pitfalls. You need to know C reasonably well before you will
succeed at exploring new algorithms using it. Java might be a good choice but
the amount of boilerplate needed to get a simple experiment running is
tedious, and the "everything is an object" restriction can impede elegant
implementations. Functional languages are awesome but they abstract over the
details of the computation so well that I don't think they are useful for the
first steps toward learning to think like a computer.

For these reasons, I think Golang is actually a stand-out choice for
exploring algorithms. It has a simple syntax that mostly gets out of your
way, it has excellent and well-documented tooling that you can use to check
your implementations (including benchmarking), and it gives you enough access
to low-level primitives to get a solid grasp on what the computer is doing.

## Implementations

So far I have implemented:

- Quicksort
- Binary trees
- Array shuffle
- Binary search with benchmarks
- Min Heap

## WIP

- Bloom filter
- Fibonnaci heap