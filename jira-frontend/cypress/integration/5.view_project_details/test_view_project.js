import { constant_values } from "../../constants";

describe("Checks the create project view", () => {

    it("should test the project create view ", () => {
      // localStorage.setItem('access-token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2NTA0OTc3NzQsImlhdCI6MTY1MDQ5NDE3NH0.i35ZnXwfDfLNhu2oiCONB2cohHs3uoUCwCH9GPdDot4')
      localStorage.setItem('access-token', `${constant_values.ACCESS_TOKEN}`)

      
    //   cy.get("[id=create_project]").click();
      cy.contains('Project 1').click();
      cy.wait(2000);
    //   cy.get("[id=proj_name]").type("sample_proj");
    //   cy.get("[id=proj_desc]").type("sample_decription");
    //   cy.get("[id=proj_key]").type("sample_proj_key");
  
    //   cy.get("[id=create_proj_btn]").click();
    //   cy.visit(`${constant_values.PROJECT_LIST}`);
    
    });
  
  });