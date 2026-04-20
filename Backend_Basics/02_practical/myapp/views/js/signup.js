function signUp() {
  let data = getFormData();
  console.log(data);

  if (data.p1 != data.p2) {
    alert("The password did not match");
    return;
  } else if (data.email == "") {
    alert("Email cannot be empty bruh");
    return;
  } else if (data.firstname == "") {
    alert("Firstname cannot be empty, so sorry");
    return;
  }

  data.password = data.p1;
  delete data.p2;
  delete data.p1;

  console.log(data);
  fetch("/signup", {
    method: "POST",
    body: JSON.stringify(data),
  }).then((res) => {
    if (res.ok) {
      console.log("Sign up sucessful");
    } else {
      console.log("Sign up error", res);
    }
  });
}

function getFormData() {
  console.log("Getting form data");
  return {
    firstname: document.getElementById("fname").value,
    lastname: document.getElementById("lname").value,
    email: document.getElementById("email").value,
    p1: document.getElementById("pw1").value,
    p2: document.getElementById("pw2").value,
  };
}
