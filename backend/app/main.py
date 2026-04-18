from fastapi import FastAPI
from app.db.database import engine
from app.db.base import Base
from app.models import product

app = FastAPI()

Base.metadata.create_all(bind=engine)