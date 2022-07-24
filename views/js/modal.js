// Get DOM Elements
const modal = document.querySelector("#my-modal");
const modalBtn = document.querySelectorAll(".modal-btn");
const closeBtn = document.querySelector(".close");

console.log(modalBtn);

// Events
// modalBtn.map((modalButton) => modalButton.addEventListener("click", openModal));
modalBtn.forEach((modalButton) => {
  modalButton.addEventListener("click", openModal);
});
// modalBtn.addEventListener("click", openModal);
closeBtn.addEventListener("click", closeModal);
window.addEventListener("click", outsideClick);

// Open
function openModal() {
  modal.style.display = "block";
  modal.style.backgroundColor = "rgba(0, 0, 0, 0.25)";
  modal.style.zIndex = "10000";
  modal.style.backdropFilter = "blur(25px)";
}

// Close
function closeModal() {
  modal.style.display = "none";
}

// Close If Outside Click
function outsideClick(e) {
  if (e.target == modal) {
    modal.style.display = "none";
  }
}
