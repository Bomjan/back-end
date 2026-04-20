function login() {
  let data = getLoginForm();

  fetch("/login", {
    method: "POST",
    body: JSON.stringify(data),
    headers: { "Content-Type": "application/json; charset=UTF-8" },
    credentials: "include",
  })
    .then((res) => res.text())
    .then((data) => {
      let admin = JSON.parse(data);
      console.log(
        `Welcome ${admin.firstname} ${admin.lastname}, we will update on ${admin.email}`,
      );
    });
}

const getLoginForm = () => {
  return {
    email: document.getElementById("email").value,
    password: document.getElementById("pw").value,
  };
};

function logout() {
  fetch("/logout", {
    credentials: "include",
  })
    .then((res) => {
      if (res.ok) {
        window.open("index.html", "_self");
      } else {
        throw new Error(res.statusText);
      }
    })
    .catch((e) => {
      alert(e);
    });
}
