import { ComponentFixture, TestBed } from '@angular/core/testing';
import { NgbDropdown, NgbDropdownModule } from '@ng-bootstrap/ng-bootstrap';

import { NavbarComponent } from './navbar.component';

describe('NavbarComponent', () => {
  let component: NavbarComponent;
  let fixture: ComponentFixture<NavbarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NavbarComponent ],
      imports:[NgbDropdownModule]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create navbar component', () => {
    expect(component).toBeTruthy();
  });
  it('should have',()=>{
    const dropdown = fixture.debugElement.nativeElement.querySelector('.d-inline-block');
    expect(dropdown.innerHTML).not.toBeNull();
    const b = fixture.debugElement.nativeElement.querySelector('.example-icon');
    expect(b.innerHTML).not.toBeNull();
    
  })

});
