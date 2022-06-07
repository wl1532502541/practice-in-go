// var ref = require("ref");
var ffi = require("ffi-napi");
// var Struct = require("ref-struct");
// var ArrayType = require("ref-array");

// var longlong = ref.types.longlong;
// var LongArray = ArrayType(longlong);

// define object GoSlice to map to:
// C type struct { void *data; GoInt len; GoInt cap; }
// var GoSlice = Struct({
//   data: LongArray,
//   len: "longlong",
//   cap: "longlong",
// });

// define object GoString to map:
// C type struct { const char *p; GoInt n; }
// var GoString = Struct({
//   p: "string",
//   n: "longlong",
// });

// define foreign functions
var goFuncs = ffi.Library("./go-funcs.dylib", {
  Add: ["longlong", ["longlong", "longlong"]],
  Cosine: ["double", ["double"]],
  // Sort: ["void", [GoSlice]],
  //   Log: ["longlong", [GoString]],
});

// call Add
console.log("goFuncs.Add(12, 99) = ", goFuncs.Add(12, 99));

// call Cosine
console.log("goFuncs.Cosine(1) = ", goFuncs.Cosine(1));

// call Sort
// nums = LongArray([12, 54, 0, 423, 9]);
// var slice = new GoSlice();
// slice["data"] = nums;
// slice["len"] = 5;
// slice["cap"] = 5;
// goFuncs.Sort(slice);
// console.log("goFuncs.Sort([12,54,9,423,9] = ", nums.toArray());

// call Log
// str = new GoString();
// str["p"] = "Hello Node!";
// str["n"] = 11;
// goFuncs.Log(str);
