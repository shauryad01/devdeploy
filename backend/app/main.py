from fastapi import FastAPI, Depends
from app.db.database import engine, get_db
from app.db.base import Base
from app.models.product import Product
from sqlalchemy.orm import Session
from app.schemas.product import ProductCreate, ProductOut

app = FastAPI()

Base.metadata.create_all(bind=engine)

@app.get("/test-db")
def test_db(db:Session = Depends(get_db)):
    products = db.query(Product).all()
    return products

@app.post("/products", response_model=ProductOut)
def add_new_product(product: ProductCreate, db:Session = Depends(get_db), ):
    new_prod = Product(product.name, price = product.price)
    db.add(new_prod)
    db.commit()
    db.refresh(new_prod)

    return new_prod
