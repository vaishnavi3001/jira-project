import { constant_values } from "../../constants";

describe("Checks the sprint-detail view", () => {

    it("should test if he sprint detail view opens", () => {
    
      localStorage.setItem('access-token', `${constant_values.ACCESS_TOKEN}`)

      cy.contains('Go to sprint').click();
      cy.wait(2000);
    
    });
  
  });