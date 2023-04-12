

function generateKey() {
  $.ajax({
    method: "POST",
    url: "/",
    data: {
      data_action: "GENERATE",
    },
    success: function () {
      console.log("KEY GENERATED")
    },
  });
}

// const characters =
//   "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";

// function generateString(length) {
//   let result = "";
//   const charactersLength = characters.length;
//   for (let i = 0; i < length; i++) {
//     result += characters.charAt(Math.floor(Math.random() * charactersLength));
//   }

//   return result;
// }

