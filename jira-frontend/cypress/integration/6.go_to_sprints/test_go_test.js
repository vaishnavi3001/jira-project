import { constant_values } from "../../constants";

describe("Checks the sprint view", () => {

    it("should test if he sprint view opens", () => {
      // localStorage.setItem('access-token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NTA0OTc3NzQsImlhdCI6MTY1MDQ5NDE3NH0.i35ZnXwfDfLNhu2oiCONB2cohHs3uoUCwCH9GPdDot4')
      localStorage.setItem('access-token', `${constant_values.ACCESS_TOKEN}`)

      cy.contains('Go To Sprints').click();
      cy.wait(2000);

    
    });
  
  });