# nestjson

```
Given a “flatten” dictionary object, whose keys are dot-separated.
For example, { ‘A’: 1, ‘B.A’: 2, ‘B.B’: 3, ‘CC.D.E’: 4, ‘CC.D.F’: 5}.
Implement a function in any language to transform it to a “nested” dictionary object.
In the above case, the nested version is like:
{
  ‘A’: 1,
  ‘B’: {
    ‘A’: 2,
    ‘B’: 3,
  },
  ‘CC’: {
    ‘D’: {
      ‘E’: 4,
      ‘F’: 5,
    }
  }
}   
It’s guaranteed that no keys in dictionary are prefixes of other keys.
```
