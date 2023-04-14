function generateKey() {
  $.ajax({
    method: "POST",
    url: "/",
    data: {
      data_action: "GENERATE KEY",
    },
    success: function () {
      // $("#key-input").text(result)
      console.log("KEY GENERATED")
    
    },
  });
}

function generateTOTP() {
  // $("#genSecret").text(localStorage.getItem("key"))

  $.ajax({
    method: "POST",
    url: "/",
    data: {
      data_action: "GENERATE TOTP",
    },
    success: function () {
      // $("#key-input").text(result);
    },
  });
}

