package io.davinash;


import org.apache.geode.cache.execute.Function;
import org.apache.geode.cache.execute.FunctionContext;
import org.apache.geode.cache.execute.FunctionException;
import org.apache.geode.cache.execute.RegionFunctionContext;

/**
 * Hello world!
 */
public class MyFunction implements Function {

  @Override
  public boolean hasResult() {
    return true;
  }

  @Override
  public void execute(FunctionContext fc) {
    if (!(fc instanceof RegionFunctionContext)) {
      throw new FunctionException("This is a data aware function, and has to "
          + "be called using FunctionService.onRegion.");
    }
    RegionFunctionContext context = (RegionFunctionContext) fc;

    context.getResultSender().sendResult("Result-MyFunction-Success-1");
    context.getResultSender().lastResult("Result-MyFunction-Last-Result");
  }

  @Override
  public String getId() {
    return "MyFunction";
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
