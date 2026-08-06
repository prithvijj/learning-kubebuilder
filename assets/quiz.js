// Shared quiz widget for the Kubebuilder teaching workspace.
// Markup contract:
// <div class="quiz-mcq" data-answer="b">
//   <p class="quiz-q">Question text</p>
//   <button data-choice="a">Choice A</button>
//   <button data-choice="b">Choice B</button>
//   <p class="quiz-feedback"></p>
// </div>
//
// <div class="quiz-order" data-answer="1,3,2,4">
//   <p class="quiz-q">Question text</p>
//   <ol class="quiz-order-list">
//     <li data-step="1" draggable="true">Step text</li>
//     ...
//   </ol>
//   <button class="quiz-check">Check order</button>
//   <p class="quiz-feedback"></p>
// </div>

document.addEventListener("DOMContentLoaded", () => {
  initMCQ();
  initOrder();
});

function initMCQ() {
  document.querySelectorAll(".quiz-mcq").forEach((q) => {
    const answer = q.dataset.answer;
    const feedback = q.querySelector(".quiz-feedback");
    q.querySelectorAll("button[data-choice]").forEach((btn) => {
      btn.addEventListener("click", () => {
        q.querySelectorAll("button[data-choice]").forEach((b) => {
          b.disabled = true;
          b.classList.remove("correct", "incorrect");
        });
        const correct = btn.dataset.choice === answer;
        btn.classList.add(correct ? "correct" : "incorrect");
        if (!correct) {
          const rightBtn = q.querySelector(`button[data-choice="${answer}"]`);
          if (rightBtn) rightBtn.classList.add("correct");
        }
        if (feedback) {
          feedback.textContent = correct
            ? "Correct."
            : "Not quite — the correct answer is highlighted.";
          feedback.classList.toggle("good", correct);
          feedback.classList.toggle("bad", !correct);
        }
      });
    });
  });
}

function initOrder() {
  document.querySelectorAll(".quiz-order").forEach((q) => {
    const list = q.querySelector(".quiz-order-list");
    const checkBtn = q.querySelector(".quiz-check");
    const feedback = q.querySelector(".quiz-feedback");
    let dragEl = null;

    list.querySelectorAll("li").forEach((li) => {
      li.addEventListener("dragstart", () => {
        dragEl = li;
        li.classList.add("dragging");
      });
      li.addEventListener("dragend", () => li.classList.remove("dragging"));
      li.addEventListener("dragover", (e) => {
        e.preventDefault();
        const after = getDragAfterElement(list, e.clientY);
        if (after == null) list.appendChild(dragEl);
        else list.insertBefore(dragEl, after);
      });
    });

    checkBtn.addEventListener("click", () => {
      const answer = q.dataset.answer.split(",");
      const current = Array.from(list.querySelectorAll("li")).map(
        (li) => li.dataset.step
      );
      const correct = JSON.stringify(current) === JSON.stringify(answer);
      list.querySelectorAll("li").forEach((li, i) => {
        li.classList.toggle("correct", li.dataset.step === answer[i]);
        li.classList.toggle("incorrect", li.dataset.step !== answer[i]);
      });
      if (feedback) {
        feedback.textContent = correct
          ? "Correct order."
          : "Not quite — items in red are out of place. Drag to reorder and check again.";
        feedback.classList.toggle("good", correct);
        feedback.classList.toggle("bad", !correct);
      }
    });
  });
}

function getDragAfterElement(container, y) {
  const els = [...container.querySelectorAll("li:not(.dragging)")];
  return els.reduce(
    (closest, child) => {
      const box = child.getBoundingClientRect();
      const offset = y - box.top - box.height / 2;
      if (offset < 0 && offset > closest.offset) {
        return { offset, element: child };
      }
      return closest;
    },
    { offset: Number.NEGATIVE_INFINITY, element: null }
  ).element;
}
