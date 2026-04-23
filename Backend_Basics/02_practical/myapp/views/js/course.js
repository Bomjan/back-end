const $ = (id) => document.getElementById(id);
const getFormData = () => {
  return {
    courseId: parseInt($("cid").value),
    courseName: $("cname").value,
  };
};
const resetForm = () => {
  document.getElementById("cid").value = "";
  document.getElementById("cname").value = "";
};

window.onload = getCourse();
function addCourse() {
  let data = getFormData();

  fetch("/course/add", {
    method: "POST",
    body: JSON.stringify(data),
    headers: { "Content-Type": "application/json; charset=UTF-8" },
  }).then((res) => {
    if (res.ok) {
      console.log("Course Added sucessfully");
      showCourse(data);
    } else {
      console.log("Some errors in adding the course", res);
    }

    resetForm();
  });
}

function getCourse() {
  fetch("/course")
    .then((res) => res.text())
    .then((data) => JSON.parse(data))
    .then((obj) => obj.forEach((c) => showCourse(c)));
}

function showCourse(course) {
  let table = $("myTable");
  // table.innerHTML = "";
  row = table.insertRow(table.length);
  td = [];

  for (let i = 0; i < table.rows[0].cells.length; i++) {
    td[i] = row.insertCell(i);
  }

  td[0].innerHTML = course.courseId;
  td[1].innerHTML = course.courseName;
  td[2].innerHTML = `<input type="button" onclick="deleteCourse(this)" value="delete" id="button-1">`;
  td[3].innerHTML = `<input type="button" onclick="updateCourse(this)" value="update" id="button-add">`;
}

function deleteCourse(r) {
  const cid = r.parentElement.parentElement.cells[0].textContent;
  fetch("/course/" + cid, {
    method: "DELETE",
    headers: { "Content-Type": "application/json; charset=UTF-8" },
  }).then((res) => {
    if (res.ok) {
      // also remove from the table
      resetTable();
      getCourse();
    }
  });
}
function updateCourse(r) {
  const selectedRow = r.parentElement.parentElement;
  const cid = selectedRow.cells[0].textContent;
  const oldData = {
    courseId: parseInt(cid),
    courseName: selectedRow.cells[1].textContent,
  };
  updateForm(cid, oldData);
}

const resetTable = () =>
  ($("myTable").innerHTML = `<tr>
      <th>Course ID</th>
      <th>Course Name</th>
      <th></th>
      <th></th>
  </tr>`);

function updateForm(cid, oldData) {
  const button = $("button-add");
  button.textContent = "Update";
  button.setAttribute("onclick", `update(${cid})`);
  document.getElementById("cid").value = oldData.courseId;
  document.getElementById("cname").value = oldData.courseName;
}

function update(cid) {
  const newData = getFormData();
  fetch("/course/" + cid, {
    method: "PUT",
    body: JSON.stringify(newData),
    headers: { "Content-Type": "application/json; charset=UTF-8" },
  }).then((res) => {
    if (res.ok) {
      console.log("Course updated successfully");
      resetTable();
      getCourse();
      resetForm();
      $("button-add").textContent = "Add";
      $("button-add").setAttribute("onclick", "addCourse()");
    } else {
      alert("Error updating course");
    }
  });
}
