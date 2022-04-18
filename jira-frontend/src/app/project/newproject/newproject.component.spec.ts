import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { NewprojectComponent } from './newproject.component';

describe('NewprojectComponent', () => {
  let component: NewprojectComponent;
  let fixture: ComponentFixture<NewprojectComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewprojectComponent ],
      imports:[HttpClientModule]
      
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewprojectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should have',()=>{
    const title = fixture.debugElement.nativeElement.querySelector('.modal-title');
    expect(title.innerHTML).not.toBeNull();
    const footer = fixture.debugElement.nativeElement.querySelector('.modal-footer');
    expect(footer.innerHTML).not.toBeNull();
    
  })
});
