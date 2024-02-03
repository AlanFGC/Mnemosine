package org.alanfgc.representations;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.util.List;

public class Answer {
    private int field;
    private List<String> answers;
    private List<String> incorrectAnswers;
    private String explanation;
    private QuestionType questionType;

    public Answer() {
        // Jackson deserialization
    }

    @JsonProperty
    public int getField() {
        return field;
    }

    @JsonProperty
    public List<String> getAnswers() {
        return answers;
    }

    @JsonProperty
    public List<String> getIncorrectAnswers() {
        return incorrectAnswers;
    }

    @JsonProperty
    public String getExplanation() {
        return explanation;
    }

    @JsonProperty
    public QuestionType getQuestionType() {
        return questionType;
    }

    @JsonProperty
    public void setField(int field) {
        this.field = field;
    }

    @JsonProperty
    public void setAnswers(List<String> answers) {
        this.answers = answers;
    }

    @JsonProperty
    public void setIncorrectAnswers(List<String> incorrectAnswers) {
        this.incorrectAnswers = incorrectAnswers;
    }

    @JsonProperty
    public void setExplanation(String explanation) {
        this.explanation = explanation;
    }

    @JsonProperty
    public void setQuestionType(QuestionType questionType) {
        this.questionType = questionType;
    }
}