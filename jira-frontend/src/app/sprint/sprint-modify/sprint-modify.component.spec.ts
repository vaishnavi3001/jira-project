import { HttpClient, HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SprintModifyComponent } from './sprint-modify.component';

describe('SprintModifyComponent', () => {
  let component: SprintModifyComponent;
  let fixture: ComponentFixture<SprintModifyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SprintModifyComponent ],
      imports: [ HttpClientModule ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SprintModifyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

   it('should have',()=>{
   const header = fixture.debugElement.nativeElement.querySelector('.card-header');
    expect(header.innerHTML).not.toBeNull();
    const body = fixture.debugElement.nativeElement.querySelector('.card-body');
    expect(body.innerHTML).not.toBeNull();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).toContain('Edit Sprint');
    
  })
});
