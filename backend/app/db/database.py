from sqlalchemy import orm, create_engine
from dotenv import load_dotenv
import os
from sqlalchemy.orm import sessionmaker
from pathlib import Path

env_path = Path(__file__).resolve().parent.parent.parent / ".env"
load_dotenv(dotenv_path=env_path)
DB_URL = os.getenv("DATABASE_URL")
if not DB_URL:
    raise RuntimeError("DATABASE_URL is not set")

engine = create_engine(DB_URL,echo=True)
SessionLocal= sessionmaker(autocommit=False, autoflush=False, bind=engine)

def get_db():
    db=SessionLocal()
    try:
        yield db
    finally:
        db.close()