package org.alanfgc;

import io.dropwizard.core.Application;
import io.dropwizard.core.setup.Environment;
import org.alanfgc.Resources.FlashCardResource;

public class DedalosApplicationStart extends Application<DedalosConfiguration> {
    public static void main(String[] args) throws Exception {
        System.out.println("Dedalos Starting...");
        new DedalosApplicationStart().run(args);

    }

    @Override
    public void run(DedalosConfiguration dedalosConfiguration, Environment environment) throws Exception {
        FlashCardResource flashcardResource = new FlashCardResource();
        environment.jersey().register(flashcardResource);
    }
}
