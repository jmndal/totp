function generateKey() {
  $.ajax({
    method: "POST",
    url: "/",
    data: {
      data_action: "GENERATE_KEY",
    },
    success: function () {
      // $("#key-input").text(result)
      console.log("KEY GENERATED")
    },
  });
}

function generateTOTP() {
  $.ajax({
    method: "POST",
    url: "/",
    data: {
      data_action: "GENERATE_TOTP",
    },
    success: function () {
      // $("#key-input").text(result);
      console.log("TOTP GENERATED")
    },
  });
}

