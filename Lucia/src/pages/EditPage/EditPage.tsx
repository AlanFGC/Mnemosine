import { useParams } from 'react-router-dom';
import FlashCardEditor from '../../components/FlashCardEditor/FlashCardEditor';

export default function EditPage() {
  const { id } = useParams();
  if (id) {
    return (
      <div>
        <FlashCardEditor flashcardID={id} />
      </div>
    );
  }

  return (
    <div>
      <FlashCardEditor flashcardID={undefined} />
    </div>
  );
}
