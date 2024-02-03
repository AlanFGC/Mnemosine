package org.alanfgc.Resources;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;

@Path("/flashcard")
public class FlashCardResource {

    @GET
    @Produces(MediaType.TEXT_PLAIN)
    public String Hello(){
        return "Hello World";
    }



}
