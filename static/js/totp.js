

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

function generateTOTP() {
  $.ajax({
    method: "POST",
    url: "/",
    data: {
      data_action: "GENERATE TOTP",
    },
    success: function () {
      console.log("KEY")
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

