import { faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Nutrition from "./Nutrition";

function Modal({ nutrition, setIsOpen }) {
  return (
    <div className="bg-[#1515153b]">
      <div className="bg-white drop-shadow-lg rounded-lg m-4 p-4 w-[360px] h-fit max-h-[90vh] overflow-y-auto absolute inset-0">
        <div className="flex items-center justify-between mb-3">
          <h1
            className="text-xl font-bold"
            style={{ color: nutrition.grade_detail.color }}
          >
            {nutrition.grade_detail.title}
          </h1>
          <FontAwesomeIcon
            icon={faXmark}
            onClick={() => setIsOpen(false)}
            className="cursor-pointer"
          />
        </div>

        <p className="text-xs">{nutrition.grade_detail.description}</p>

        <a
          className="text-[10px] underline text-[#0075FF]"
          href={nutrition.grade_detail.url}
        >
          Here detail about Nutri-Score
        </a>

        <img
          className="object-cover mt-1 w-72 h-44"
          src="/grade.png"
          alt="nutri-score"
        />

        <div className="mt-4">
          <h2 className="font-bold text-sm mb-2">Ingredients:</h2>
          <ul className="text-xs list-disc list-inside mb-4 text-[var(--grey)]">
            {nutrition.ingredients?.map((ing, idx) => (
              <li key={idx}>{ing.name} ({ing.measurements})</li>
            ))}
          </ul>
          
          <h2 className="font-bold text-sm mb-2">Nutrition Facts:</h2>
          <Nutrition nutrition={nutrition.nutrition} />
        </div>
      </div>
    </div>
  );
}

export default Modal;
