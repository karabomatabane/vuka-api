import { ComponentFixture, TestBed } from '@angular/core/testing';

import { XSidenavComponent } from './x-sidenav.component';

describe('XSidenavComponent', () => {
  let component: XSidenavComponent;
  let fixture: ComponentFixture<XSidenavComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [XSidenavComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(XSidenavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
