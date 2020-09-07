package io.davinash;


import org.apache.geode.cache.execute.Function;
import org.apache.geode.cache.execute.FunctionContext;
import org.apache.geode.cache.execute.FunctionException;
import org.apache.geode.cache.execute.RegionFunctionContext;

/**
 * MyFunctionException!
 */
public class MyFunctionException implements Function {

  @Override
  public boolean hasResult() {
    return true;
  }

  @Override
  public void execute(FunctionContext fc) {
       throw new FunctionException("This is Dummy Exception");

  }

  @Override
  public String getId() {
    return "MyFunctionException";
  }

  @Override
  public boolean optimizeForWrite() {
    return false;
  }

  @Override
  public boolean isHA() {
    return false;
  }

}
