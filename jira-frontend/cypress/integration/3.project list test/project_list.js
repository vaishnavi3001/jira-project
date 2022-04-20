import { constant_values } from "../../constants";

describe("Checks project list view", () => {

    it("should test the project list view page", () => {
      localStorage.setItem('access-token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NTA0OTc3NzQsImlhdCI6MTY1MDQ5NDE3NH0.i35ZnXwfDfLNhu2oiCONB2cohHs3uoUCwCH9GPdDot4')

      cy.visit(`${constant_values.PROJECT_LIST}`);
      
    //   cy.get("[id=email]").type("pypalkar23");
    //   cy.get("[id=password]").type("mandar");
  
      // cy.get("[id=login-button]").click();
    
    });
  
  });