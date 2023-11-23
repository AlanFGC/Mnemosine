import AnswerEditor from '../../components/AnswerEditor/AnswerEditor';
import FlashCardEditor from '../../components/FlashCardEditor/FlashCardEditor';

export default function EditPage() {
  return (
    <div className="App">
      <FlashCardEditor />
      <AnswerEditor field={0} />
    </div>
  );
}
