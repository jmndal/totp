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
    
      console.log("hi")
     

    },
  });
}

function validateOTP() {
  // Send a POST request to the server to validate the OTP
  console.log("test", $("#otp-input").text() == $("#totp").val())
  if ($("#otp-input").text() == $("#totp").val()) {
    $("#otp-status").text("TOTP code is valid!") 
  } else {
    $("#otp-status").text("Invalid TOTP code!") 
  }
}

function updateTimer() {
  var now = new Date();
  var timeLeft = 30 - now.getSeconds() % 30;
  $("#timer").text(timeLeft + ' seconds');

  if (now.getSeconds() % 30 === 0) {
    // Wait 1 second and reload the page
    setTimeout(function () {
      location.reload();
    }, 5);
  }
}

console.log("hi")
setInterval(updateTimer, 1000);
