package org.alanfgc.representations;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.time.LocalDateTime;
import java.util.List;

public class UserFlashCard {
    private String id;
    private String username;
    private String title;
    private String text;
    private List<Answer> answers;
    private List<String> media;
    private List<String> languages;
    private List<String> topics;
    private LocalDateTime dateCreated;

    public UserFlashCard() {
    }
    public UserFlashCard(
            String id, String username, String title, String text, List<Answer> answers,
            List<String> media, List<String> languages, List<String> topics, LocalDateTime dateCreated
    ) {
        this.id = id;
        this.username = username;
        this.title = title;
        this.text = text;
        this.answers = answers;
        this.media = media;
        this.languages = languages;
        this.topics = topics;
        this.dateCreated = dateCreated;
    }


    @JsonProperty
    public String getId() {
        return id;
    }

    @JsonProperty
    public void setId(String id) {
        this.id = id;
    }

    @JsonProperty
    public String getUsername() {
        return username;
    }

    @JsonProperty
    public void setUsername(String username) {
        this.username = username;
    }

    @JsonProperty
    public String getTitle() {
        return title;
    }

    @JsonProperty
    public void setTitle(String title) {
        this.title = title;
    }

    @JsonProperty
    public String getText() {
        return text;
    }
    @JsonProperty
    public void setText(String text) {
        this.text = text;
    }

    @JsonProperty
    public List<Answer> getAnswers() {
        return answers;
    }

    @JsonProperty
    public void setAnswers(List<Answer> answers) {
        this.answers = answers;
    }

    @JsonProperty
    public List<String> getMedia() {
        return media;
    }

    @JsonProperty
    public void setMedia(List<String> media) {
        this.media = media;
    }

    @JsonProperty
    public List<String> getLanguages() {
        return languages;
    }

    @JsonProperty
    public void setLanguages(List<String> languages) {
        this.languages = languages;
    }

    @JsonProperty
    public List<String> getTopics() {
        return topics;
    }

    @JsonProperty
    public void setTopics(List<String> topics) {
        this.topics = topics;
    }

    @JsonProperty
    public LocalDateTime getDateCreated() {
        return dateCreated;
    }

    @JsonProperty
    public void setDateCreated(LocalDateTime dateCreated) {
        this.dateCreated = dateCreated;
    }
}
