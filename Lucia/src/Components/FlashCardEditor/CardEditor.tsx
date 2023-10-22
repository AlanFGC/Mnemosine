import { useState, useCallback } from 'react';
import ReactQuill from 'react-quill';
import 'react-quill/dist/quill.snow.css';

interface CardEditorProps {
  onTextChange: (text: string) => void;
}

function CardEditor({ onTextChange }: CardEditorProps) {
  const [value, setValue] = useState('');

  const handleTextChange = useCallback((text: string) => {
    setValue(text);
    onTextChange(text);
  }, [onTextChange]);

  return <ReactQuill theme="snow" value={value} onChange={handleTextChange} />;
}

export default CardEditor;
