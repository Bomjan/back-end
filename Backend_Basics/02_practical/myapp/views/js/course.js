const $ = (id) => document.getElementById(id);
const getFormData = () => {
  return {
    courseId: parseInt($("cid").value),
    courseName: $("cname").value,
  };
};

function addCourse() {
  let data = getFormData();

  fetch("/course/add", {
    method: "POST",
    body: JSON.stringify(data),
    headers: { "Content-Type": "application/json; charset=UTF-8" },
  }).then((res) => {
    if (res.ok) {
      console.log("Course Added sucessfully");
    } else {
      console.log("Some errors in adding the course", res);
    }
  });
}
