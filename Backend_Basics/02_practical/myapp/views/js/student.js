function addStudent() {
  var data = getFormData();

  var sid = data.stdId;

  if (isNaN(sid)) {
    alert("Enter valid student id");
    return;
  } else if (data.email == "") {
    alert("Email cannot be empty");
    return;
  } else if (data.fname == "") {
    alert("First name cannot be empty");
  }

  console.log("Sending data:", data);

  console.log(data);

  fetch("/student/add", {
    method: "POST",
    body: JSON.stringify(data),
    headers: { "Content-Type": "application/json; charset=UTF-8" },
    credentials: "include",
  })
    .then((res1) => {
      if (res1.ok) {
        fetch("/student/" + sid, {
          credentials: "include",
        })
          .then((res2) => res2.text())
          .then((dat) => showStudent(dat));
      } else {
        throw new Error(res1.status);
      }
    })
    .catch((e) => {
      if (e.message == 303) {
        alert("User not logged in");
        window.open("index.html", "_self");
      } else if (e.message == 500) {
        alert("Internal Server Error");
      } else if (e.message == 401) {
        alert("You are not authorized");
      }
    });

  resetForm();
}

function showStudent(data) {
  const student = JSON.parse(data);
  var table = document.getElementById("myTable");
  var row = table.insertRow(table.length);
  var td = [];

  for (let i = 0; i < table.rows[0].cells.length; i++) {
    td[i] = row.insertCell(i);
  }
  td[0].innerHTML = student.stdId;
  td[1].innerHTML = student.firstName;
  td[2].innerHTML = student.lastName;
  td[3].innerHTML = student.email;
  td[4].innerHTML = `<input type="button" onclick="deleteStudent(this)" value="delete" id="button-1">`;
  td[5].innerHTML = `<input type="button" onclick="updateStudent(this)" value="update" id="button-add">`;
}

function resetForm() {
  document.getElementById("sid").value = "";
  document.getElementById("fname").value = "";
  document.getElementById("lname").value = "";
  document.getElementById("email").value = "";
}

window.onload = function () {
  fetch("/student", {
    credentials: "include",
  })
    .then((res) => res.text())
    .then((data) => {
      data = JSON.parse(data);
      for (let i of data) {
        showStudent(JSON.stringify(i));
      }
      // showStudent(data);
    });
};

function updateStudent(r) {
  selectedRow = r.parentElement.parentElement;
  document.getElementById("sid").value = selectedRow.cells[0].innerHTML;
  document.getElementById("fname").value = selectedRow.cells[1].innerHTML;
  document.getElementById("lname").value = selectedRow.cells[2].innerHTML;
  document.getElementById("email").value = selectedRow.cells[3].innerHTML;

  var btn = document.getElementById("button-add");
  sid = selectedRow.cells[0].innerHTML;

  if (btn) {
    btn.innerHTML = "Update";
    btn.setAttribute("onclick", "update(sid)");
  }
}

function update(sid) {
  var data = getFormData();
  fetch("/student/" + sid, {
    method: "PUT",
    body: JSON.stringify(data),
    headers: { "Content-Type": "application/json; charset=UTF-8" },
    credentials: "include",
  }).then((res) => {
    if (res.ok) {
      selectedRow.cells[0].innerHTML = data.stdId;
      selectedRow.cells[1].innerHTML = data.firstName;
      selectedRow.cells[2].innerHTML = data.lastName;
      selectedRow.cells[3].innerHTML = data.email;

      var button = document.getElementById("button-add");
      button.innerHTML = "Add";
      button.setAttribute("onclick", "addStudent()");
      selectedRow = null;
      resetForm();
    } else {
      console.log(res);
      alert("Server: Update error");
    }
  });
}

function getFormData() {
  return {
    stdId: parseInt(document.getElementById("sid").value),
    firstName: document.getElementById("fname").value,
    lastName: document.getElementById("lname").value,
    email: document.getElementById("email").value,
  };
}

function deleteStudent(r) {
  if (confirm("Are you sure you want to delete this?")) {
    selectedRow = r.parentElement.parentElement;
    sid = selectedRow.cells[0].innerHTML;

    fetch("/student/" + sid, {
      method: "DELETE",
      headers: { "Content-Type": "application/json; charset=UTF-8" },
      credentials: "include",
    }).then((res) => {
      if (res.ok) {
        console.log("Deleted sucessfylly");
      } else {
        console.log("You sucked");
      }
    });

    var rowIndex = selectedRow.rowIndex;

    if (rowIndex > 0) {
      document.getElementById("myTable").deleteRow(rowIndex);
    }

    selectedRow = null;
  }
}
